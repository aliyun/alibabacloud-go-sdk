// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityIdentifyResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSecurityIdentifyResultsResponseBody
	GetCode() *string
	SetData(v bool) *DeleteSecurityIdentifyResultsResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteSecurityIdentifyResultsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteSecurityIdentifyResultsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSecurityIdentifyResultsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSecurityIdentifyResultsResponseBody
	GetSuccess() *bool
}

type DeleteSecurityIdentifyResultsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSecurityIdentifyResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityIdentifyResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIdentifyResultsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSecurityIdentifyResultsResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteSecurityIdentifyResultsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteSecurityIdentifyResultsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSecurityIdentifyResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityIdentifyResultsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSecurityIdentifyResultsResponseBody) SetCode(v string) *DeleteSecurityIdentifyResultsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecurityIdentifyResultsResponseBody) SetData(v bool) *DeleteSecurityIdentifyResultsResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSecurityIdentifyResultsResponseBody) SetHttpStatusCode(v int32) *DeleteSecurityIdentifyResultsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSecurityIdentifyResultsResponseBody) SetMessage(v string) *DeleteSecurityIdentifyResultsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecurityIdentifyResultsResponseBody) SetRequestId(v string) *DeleteSecurityIdentifyResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityIdentifyResultsResponseBody) SetSuccess(v bool) *DeleteSecurityIdentifyResultsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSecurityIdentifyResultsResponseBody) Validate() error {
	return dara.Validate(s)
}
