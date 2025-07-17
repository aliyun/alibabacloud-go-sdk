// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterAllUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetClusterAllUrlResponseBody
	GetCode() *int32
	SetData(v string) *GetClusterAllUrlResponseBody
	GetData() *string
	SetMessage(v string) *GetClusterAllUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetClusterAllUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetClusterAllUrlResponseBody
	GetSuccess() *bool
}

type GetClusterAllUrlResponseBody struct {
	// Status code. 200 means success, other status codes are exceptions.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters in the JSON format.
	//
	// example:
	//
	// {
	//
	//   "RequestId": "41877338-646B-5DD6-BFBE-F89F1E7245AD",
	//
	//   "Data": "{\\"clusterType\\":\\"ManagedKubernetes\\",\\"remoteWriteUrl\\":\\"http:/" }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The prompt information of the returned result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40B10E04-81E8-4643-970D-F1B38F2E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the query is successful:
	//
	// - true: success
	//
	// - false: failure
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetClusterAllUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterAllUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterAllUrlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetClusterAllUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *GetClusterAllUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetClusterAllUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterAllUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetClusterAllUrlResponseBody) SetCode(v int32) *GetClusterAllUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetClusterAllUrlResponseBody) SetData(v string) *GetClusterAllUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetClusterAllUrlResponseBody) SetMessage(v string) *GetClusterAllUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetClusterAllUrlResponseBody) SetRequestId(v string) *GetClusterAllUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterAllUrlResponseBody) SetSuccess(v bool) *GetClusterAllUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetClusterAllUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
