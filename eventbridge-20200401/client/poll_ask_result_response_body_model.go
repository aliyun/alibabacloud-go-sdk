// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPollAskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PollAskResultResponseBody
	GetCode() *string
	SetData(v *AskLumaResult) *PollAskResultResponseBody
	GetData() *AskLumaResult
	SetMessage(v string) *PollAskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *PollAskResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PollAskResultResponseBody
	GetSuccess() *bool
}

type PollAskResultResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *AskLumaResult `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Agent with name \\"xxx\\" not found for account 1186xxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A995F07C-E503-XXXXXX-9CECA8566876
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PollAskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PollAskResultResponseBody) GoString() string {
	return s.String()
}

func (s *PollAskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *PollAskResultResponseBody) GetData() *AskLumaResult {
	return s.Data
}

func (s *PollAskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PollAskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PollAskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PollAskResultResponseBody) SetCode(v string) *PollAskResultResponseBody {
	s.Code = &v
	return s
}

func (s *PollAskResultResponseBody) SetData(v *AskLumaResult) *PollAskResultResponseBody {
	s.Data = v
	return s
}

func (s *PollAskResultResponseBody) SetMessage(v string) *PollAskResultResponseBody {
	s.Message = &v
	return s
}

func (s *PollAskResultResponseBody) SetRequestId(v string) *PollAskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *PollAskResultResponseBody) SetSuccess(v bool) *PollAskResultResponseBody {
	s.Success = &v
	return s
}

func (s *PollAskResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
