// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AbortCampaignResponseBody
	GetCode() *string
	SetData(v bool) *AbortCampaignResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *AbortCampaignResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AbortCampaignResponseBody
	GetMessage() *string
	SetParams(v []*string) *AbortCampaignResponseBody
	GetParams() []*string
	SetRequestId(v string) *AbortCampaignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AbortCampaignResponseBody
	GetSuccess() *bool
}

type AbortCampaignResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the operation.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AbortCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbortCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *AbortCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *AbortCampaignResponseBody) GetData() *bool {
	return s.Data
}

func (s *AbortCampaignResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AbortCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AbortCampaignResponseBody) GetParams() []*string {
	return s.Params
}

func (s *AbortCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbortCampaignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AbortCampaignResponseBody) SetCode(v string) *AbortCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *AbortCampaignResponseBody) SetData(v bool) *AbortCampaignResponseBody {
	s.Data = &v
	return s
}

func (s *AbortCampaignResponseBody) SetHttpStatusCode(v int32) *AbortCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AbortCampaignResponseBody) SetMessage(v string) *AbortCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *AbortCampaignResponseBody) SetParams(v []*string) *AbortCampaignResponseBody {
	s.Params = v
	return s
}

func (s *AbortCampaignResponseBody) SetRequestId(v string) *AbortCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortCampaignResponseBody) SetSuccess(v bool) *AbortCampaignResponseBody {
	s.Success = &v
	return s
}

func (s *AbortCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
