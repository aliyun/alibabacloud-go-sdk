// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PauseCampaignResponseBody
	GetCode() *string
	SetData(v bool) *PauseCampaignResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *PauseCampaignResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PauseCampaignResponseBody
	GetMessage() *string
	SetParams(v []*string) *PauseCampaignResponseBody
	GetParams() []*string
	SetRequestId(v string) *PauseCampaignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PauseCampaignResponseBody
	GetSuccess() *bool
}

type PauseCampaignResponseBody struct {
	// The result code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of error message parameters.
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

func (s PauseCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *PauseCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *PauseCampaignResponseBody) GetData() *bool {
	return s.Data
}

func (s *PauseCampaignResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PauseCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PauseCampaignResponseBody) GetParams() []*string {
	return s.Params
}

func (s *PauseCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseCampaignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PauseCampaignResponseBody) SetCode(v string) *PauseCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *PauseCampaignResponseBody) SetData(v bool) *PauseCampaignResponseBody {
	s.Data = &v
	return s
}

func (s *PauseCampaignResponseBody) SetHttpStatusCode(v int32) *PauseCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PauseCampaignResponseBody) SetMessage(v string) *PauseCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *PauseCampaignResponseBody) SetParams(v []*string) *PauseCampaignResponseBody {
	s.Params = v
	return s
}

func (s *PauseCampaignResponseBody) SetRequestId(v string) *PauseCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseCampaignResponseBody) SetSuccess(v bool) *PauseCampaignResponseBody {
	s.Success = &v
	return s
}

func (s *PauseCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
