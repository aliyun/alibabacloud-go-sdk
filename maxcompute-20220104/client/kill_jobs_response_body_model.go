// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *KillJobsResponseBody
	GetData() *string
	SetHttpCode(v int32) *KillJobsResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *KillJobsResponseBody
	GetRequestId() *string
}

type KillJobsResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0abb7ede16814560741256732e91b6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s KillJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KillJobsResponseBody) GoString() string {
	return s.String()
}

func (s *KillJobsResponseBody) GetData() *string {
	return s.Data
}

func (s *KillJobsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *KillJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KillJobsResponseBody) SetData(v string) *KillJobsResponseBody {
	s.Data = &v
	return s
}

func (s *KillJobsResponseBody) SetHttpCode(v int32) *KillJobsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *KillJobsResponseBody) SetRequestId(v string) *KillJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *KillJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
