// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQwenConferenceSgTicketPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryQwenConferenceSgTicketPopResponseBody
	GetCode() *string
	SetData(v *QueryQwenConferenceSgTicketPopResponseBodyData) *QueryQwenConferenceSgTicketPopResponseBody
	GetData() *QueryQwenConferenceSgTicketPopResponseBodyData
	SetMessage(v string) *QueryQwenConferenceSgTicketPopResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryQwenConferenceSgTicketPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryQwenConferenceSgTicketPopResponseBody
	GetSuccess() *bool
}

type QueryQwenConferenceSgTicketPopResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 200
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryQwenConferenceSgTicketPopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryQwenConferenceSgTicketPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryQwenConferenceSgTicketPopResponseBody) GoString() string {
	return s.String()
}

func (s *QueryQwenConferenceSgTicketPopResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryQwenConferenceSgTicketPopResponseBody) GetData() *QueryQwenConferenceSgTicketPopResponseBodyData {
	return s.Data
}

func (s *QueryQwenConferenceSgTicketPopResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryQwenConferenceSgTicketPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryQwenConferenceSgTicketPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryQwenConferenceSgTicketPopResponseBody) SetCode(v string) *QueryQwenConferenceSgTicketPopResponseBody {
	s.Code = &v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponseBody) SetData(v *QueryQwenConferenceSgTicketPopResponseBodyData) *QueryQwenConferenceSgTicketPopResponseBody {
	s.Data = v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponseBody) SetMessage(v string) *QueryQwenConferenceSgTicketPopResponseBody {
	s.Message = &v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponseBody) SetRequestId(v string) *QueryQwenConferenceSgTicketPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponseBody) SetSuccess(v bool) *QueryQwenConferenceSgTicketPopResponseBody {
	s.Success = &v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryQwenConferenceSgTicketPopResponseBodyData struct {
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// example:
	//
	// placeholder
	ExtFields *string `json:"ExtFields,omitempty" xml:"ExtFields,omitempty"`
	// example:
	//
	// ***
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// example:
	//
	// ***
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// example:
	//
	// 23808009
	SubmitId *int64 `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
	// example:
	//
	// bPbXgB8nSzI9UIbdqAWaOhtr7T3p1Ryr
	TicketToken *string `json:"TicketToken,omitempty" xml:"TicketToken,omitempty"`
}

func (s QueryQwenConferenceSgTicketPopResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryQwenConferenceSgTicketPopResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) GetCompanyName() *string {
	return s.CompanyName
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) GetExtFields() *string {
	return s.ExtFields
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) GetFirstName() *string {
	return s.FirstName
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) GetLastName() *string {
	return s.LastName
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) GetSubmitId() *int64 {
	return s.SubmitId
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) GetTicketToken() *string {
	return s.TicketToken
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) SetCompanyName(v string) *QueryQwenConferenceSgTicketPopResponseBodyData {
	s.CompanyName = &v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) SetExtFields(v string) *QueryQwenConferenceSgTicketPopResponseBodyData {
	s.ExtFields = &v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) SetFirstName(v string) *QueryQwenConferenceSgTicketPopResponseBodyData {
	s.FirstName = &v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) SetLastName(v string) *QueryQwenConferenceSgTicketPopResponseBodyData {
	s.LastName = &v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) SetSubmitId(v int64) *QueryQwenConferenceSgTicketPopResponseBodyData {
	s.SubmitId = &v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) SetTicketToken(v string) *QueryQwenConferenceSgTicketPopResponseBodyData {
	s.TicketToken = &v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponseBodyData) Validate() error {
	return dara.Validate(s)
}
