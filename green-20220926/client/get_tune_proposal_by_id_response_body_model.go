// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTuneProposalByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTuneProposalByIdResponseBody
	GetCode() *string
	SetData(v *GetTuneProposalByIdResponseBodyData) *GetTuneProposalByIdResponseBody
	GetData() *GetTuneProposalByIdResponseBodyData
	SetMsg(v string) *GetTuneProposalByIdResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetTuneProposalByIdResponseBody
	GetRequestId() *string
}

type GetTuneProposalByIdResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTuneProposalByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTuneProposalByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTuneProposalByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTuneProposalByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTuneProposalByIdResponseBody) GetData() *GetTuneProposalByIdResponseBodyData {
	return s.Data
}

func (s *GetTuneProposalByIdResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetTuneProposalByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTuneProposalByIdResponseBody) SetCode(v string) *GetTuneProposalByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetTuneProposalByIdResponseBody) SetData(v *GetTuneProposalByIdResponseBodyData) *GetTuneProposalByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetTuneProposalByIdResponseBody) SetMsg(v string) *GetTuneProposalByIdResponseBody {
	s.Msg = &v
	return s
}

func (s *GetTuneProposalByIdResponseBody) SetRequestId(v string) *GetTuneProposalByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTuneProposalByIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTuneProposalByIdResponseBodyData struct {
	// example:
	//
	// {"example":"xxxx"}
	JsonContent *string `json:"JsonContent,omitempty" xml:"JsonContent,omitempty"`
}

func (s GetTuneProposalByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTuneProposalByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTuneProposalByIdResponseBodyData) GetJsonContent() *string {
	return s.JsonContent
}

func (s *GetTuneProposalByIdResponseBodyData) SetJsonContent(v string) *GetTuneProposalByIdResponseBodyData {
	s.JsonContent = &v
	return s
}

func (s *GetTuneProposalByIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
