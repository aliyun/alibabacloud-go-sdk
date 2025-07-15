// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIvrTrackingSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIvrTrackingSummaryResponseBody
	GetCode() *string
	SetData(v string) *GetIvrTrackingSummaryResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetIvrTrackingSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetIvrTrackingSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIvrTrackingSummaryResponseBody
	GetRequestId() *string
}

type GetIvrTrackingSummaryResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Unknown error code \\"NoPermission.Recording\\". Reason: null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DE803553-8AA9-4B9D-9E4E-A82BC69EDCEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIvrTrackingSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIvrTrackingSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetIvrTrackingSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIvrTrackingSummaryResponseBody) GetData() *string {
	return s.Data
}

func (s *GetIvrTrackingSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetIvrTrackingSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIvrTrackingSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIvrTrackingSummaryResponseBody) SetCode(v string) *GetIvrTrackingSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetIvrTrackingSummaryResponseBody) SetData(v string) *GetIvrTrackingSummaryResponseBody {
	s.Data = &v
	return s
}

func (s *GetIvrTrackingSummaryResponseBody) SetHttpStatusCode(v int32) *GetIvrTrackingSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetIvrTrackingSummaryResponseBody) SetMessage(v string) *GetIvrTrackingSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetIvrTrackingSummaryResponseBody) SetRequestId(v string) *GetIvrTrackingSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIvrTrackingSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}
