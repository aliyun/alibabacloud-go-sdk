// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAskLumaLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAskLumaLogResponseBody
	GetCode() *string
	SetData(v *QueryAskLumaLogResult) *QueryAskLumaLogResponseBody
	GetData() *QueryAskLumaLogResult
	SetMessage(v string) *QueryAskLumaLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAskLumaLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAskLumaLogResponseBody
	GetSuccess() *bool
}

type QueryAskLumaLogResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *QueryAskLumaLogResult `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AgentNotFound
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6FB52207-7621-5292-BDF2-A17E2E984160
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAskLumaLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAskLumaLogResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAskLumaLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAskLumaLogResponseBody) GetData() *QueryAskLumaLogResult {
	return s.Data
}

func (s *QueryAskLumaLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAskLumaLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAskLumaLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAskLumaLogResponseBody) SetCode(v string) *QueryAskLumaLogResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAskLumaLogResponseBody) SetData(v *QueryAskLumaLogResult) *QueryAskLumaLogResponseBody {
	s.Data = v
	return s
}

func (s *QueryAskLumaLogResponseBody) SetMessage(v string) *QueryAskLumaLogResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAskLumaLogResponseBody) SetRequestId(v string) *QueryAskLumaLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAskLumaLogResponseBody) SetSuccess(v bool) *QueryAskLumaLogResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAskLumaLogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
