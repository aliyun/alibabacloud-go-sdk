// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotspotAnalysisResponseBody
	GetCode() *string
	SetData(v string) *GetHotspotAnalysisResponseBody
	GetData() *string
	SetMessage(v string) *GetHotspotAnalysisResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotspotAnalysisResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotspotAnalysisResponseBody
	GetSuccess() *bool
}

type GetHotspotAnalysisResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHotspotAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotAnalysisResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotspotAnalysisResponseBody) GetData() *string {
	return s.Data
}

func (s *GetHotspotAnalysisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotspotAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotspotAnalysisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotspotAnalysisResponseBody) SetCode(v string) *GetHotspotAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) SetData(v string) *GetHotspotAnalysisResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) SetMessage(v string) *GetHotspotAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) SetRequestId(v string) *GetHotspotAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) SetSuccess(v bool) *GetHotspotAnalysisResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}
