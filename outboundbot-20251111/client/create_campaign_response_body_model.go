// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCampaignResponseBody
	GetCode() *string
	SetData(v string) *CreateCampaignResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateCampaignResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCampaignResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateCampaignResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateCampaignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCampaignResponseBody
	GetSuccess() *bool
}

type CreateCampaignResponseBody struct {
	// The status code of the operation.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data, which is the task ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCampaignResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateCampaignResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCampaignResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCampaignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCampaignResponseBody) SetCode(v string) *CreateCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCampaignResponseBody) SetData(v string) *CreateCampaignResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCampaignResponseBody) SetHttpStatusCode(v int32) *CreateCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCampaignResponseBody) SetMessage(v string) *CreateCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCampaignResponseBody) SetParams(v []*string) *CreateCampaignResponseBody {
	s.Params = v
	return s
}

func (s *CreateCampaignResponseBody) SetRequestId(v string) *CreateCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCampaignResponseBody) SetSuccess(v bool) *CreateCampaignResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
