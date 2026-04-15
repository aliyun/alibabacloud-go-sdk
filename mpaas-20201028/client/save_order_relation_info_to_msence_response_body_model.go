// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOrderRelationInfoToMsenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMpaasSaveOrderRelationResponse(v *SaveOrderRelationInfoToMsenceResponseBodyMpaasSaveOrderRelationResponse) *SaveOrderRelationInfoToMsenceResponseBody
	GetMpaasSaveOrderRelationResponse() *SaveOrderRelationInfoToMsenceResponseBodyMpaasSaveOrderRelationResponse
	SetRequestId(v string) *SaveOrderRelationInfoToMsenceResponseBody
	GetRequestId() *string
	SetResultCode(v string) *SaveOrderRelationInfoToMsenceResponseBody
	GetResultCode() *string
	SetResultMsg(v string) *SaveOrderRelationInfoToMsenceResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *SaveOrderRelationInfoToMsenceResponseBody
	GetSuccess() *bool
}

type SaveOrderRelationInfoToMsenceResponseBody struct {
	MpaasSaveOrderRelationResponse *SaveOrderRelationInfoToMsenceResponseBodyMpaasSaveOrderRelationResponse `json:"MpaasSaveOrderRelationResponse,omitempty" xml:"MpaasSaveOrderRelationResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// SUCCESS
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveOrderRelationInfoToMsenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveOrderRelationInfoToMsenceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveOrderRelationInfoToMsenceResponseBody) GetMpaasSaveOrderRelationResponse() *SaveOrderRelationInfoToMsenceResponseBodyMpaasSaveOrderRelationResponse {
	return s.MpaasSaveOrderRelationResponse
}

func (s *SaveOrderRelationInfoToMsenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveOrderRelationInfoToMsenceResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *SaveOrderRelationInfoToMsenceResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *SaveOrderRelationInfoToMsenceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveOrderRelationInfoToMsenceResponseBody) SetMpaasSaveOrderRelationResponse(v *SaveOrderRelationInfoToMsenceResponseBodyMpaasSaveOrderRelationResponse) *SaveOrderRelationInfoToMsenceResponseBody {
	s.MpaasSaveOrderRelationResponse = v
	return s
}

func (s *SaveOrderRelationInfoToMsenceResponseBody) SetRequestId(v string) *SaveOrderRelationInfoToMsenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceResponseBody) SetResultCode(v string) *SaveOrderRelationInfoToMsenceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceResponseBody) SetResultMsg(v string) *SaveOrderRelationInfoToMsenceResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceResponseBody) SetSuccess(v bool) *SaveOrderRelationInfoToMsenceResponseBody {
	s.Success = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceResponseBody) Validate() error {
	if s.MpaasSaveOrderRelationResponse != nil {
		if err := s.MpaasSaveOrderRelationResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveOrderRelationInfoToMsenceResponseBodyMpaasSaveOrderRelationResponse struct {
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveOrderRelationInfoToMsenceResponseBodyMpaasSaveOrderRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveOrderRelationInfoToMsenceResponseBodyMpaasSaveOrderRelationResponse) GoString() string {
	return s.String()
}

func (s *SaveOrderRelationInfoToMsenceResponseBodyMpaasSaveOrderRelationResponse) GetSuccess() *bool {
	return s.Success
}

func (s *SaveOrderRelationInfoToMsenceResponseBodyMpaasSaveOrderRelationResponse) SetSuccess(v bool) *SaveOrderRelationInfoToMsenceResponseBodyMpaasSaveOrderRelationResponse {
	s.Success = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceResponseBodyMpaasSaveOrderRelationResponse) Validate() error {
	return dara.Validate(s)
}
