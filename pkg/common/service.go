/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"context"
)

import (
	"google.golang.org/grpc"
)

// TripleGrpcService is gRPC service, used to check impl
type TripleGrpcService interface {
	ServiceDesc() *grpc.ServiceDesc
}

// TripleUnaryService is normal protocol service (except grpc service), should be implemented by users
type TripleUnaryService interface {
	InvokeWithArgs(ctx context.Context, methodName string, arguments []interface{}) (interface{}, error)
	GetReqParamsInteface(methodName string) (interface{}, bool)
}
