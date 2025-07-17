// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteAppListResponseBody
	GetCode() *int64
	SetData(v string) *DeleteAppListResponseBody
	GetData() *string
	SetMessage(v string) *DeleteAppListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAppListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAppListResponseBody
	GetSuccess() *bool
}

type DeleteAppListResponseBody struct {
	// The HTTP status code. 2XX indicates that the request was successful. 3XX indicates that the request was redirected. 4XX indicates that a request error occurred. 5XX indicates that a server error occurred.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// "{\\"code\\":200,\\"data\\":\\"{\\\\\\"code\\\\\\":200,\\\\\\"data\\\\\\":true,\\\\\\"errorCode\\\\ \\":\\\\\\"Deletion of application successful\\\\\\",\\\\\\"Message\\\\\\":\\\\\\"Deletion of application successful\\\\\\",\\\\\\"Successful\\\\\\":true,\\ \\\\"traceId\\\\\\":\\\\\\"0bc1667516940677164677396d0088\\\\\\"}\\",\\"errorCode\\":\\"Batch deletion of applications successful\\",\\"message\\":\\"Batch deletion of applications successful\\ ",\\"Success\\":true,\\"traceId\\":\\"210f470416940677161213505e0e7a\\"}"
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned for invalid request parameters.
	//
	// example:
	//
	// "Deletion of application successful"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B446DF2-3DDD-4B5B-8E3F-D5225120****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the specified applications are deleted. Valid values:
	//
	// 	- `true`: The applications are deleted.
	//
	// 	- `false`: The applications failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteAppListResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteAppListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAppListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAppListResponseBody) SetCode(v int64) *DeleteAppListResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppListResponseBody) SetData(v string) *DeleteAppListResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAppListResponseBody) SetMessage(v string) *DeleteAppListResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppListResponseBody) SetRequestId(v string) *DeleteAppListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppListResponseBody) SetSuccess(v bool) *DeleteAppListResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAppListResponseBody) Validate() error {
	return dara.Validate(s)
}
