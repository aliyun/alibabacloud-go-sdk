// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQwenConferenceSgTicketSearchPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryQwenConferenceSgTicketSearchPopResponseBody
	GetCode() *string
	SetData(v []*QueryQwenConferenceSgTicketSearchPopResponseBodyData) *QueryQwenConferenceSgTicketSearchPopResponseBody
	GetData() []*QueryQwenConferenceSgTicketSearchPopResponseBodyData
	SetMessage(v string) *QueryQwenConferenceSgTicketSearchPopResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryQwenConferenceSgTicketSearchPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryQwenConferenceSgTicketSearchPopResponseBody
	GetSuccess() *bool
}

type QueryQwenConferenceSgTicketSearchPopResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 200
	Code *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*QueryQwenConferenceSgTicketSearchPopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s QueryQwenConferenceSgTicketSearchPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryQwenConferenceSgTicketSearchPopResponseBody) GoString() string {
	return s.String()
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBody) GetData() []*QueryQwenConferenceSgTicketSearchPopResponseBodyData {
	return s.Data
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBody) SetCode(v string) *QueryQwenConferenceSgTicketSearchPopResponseBody {
	s.Code = &v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBody) SetData(v []*QueryQwenConferenceSgTicketSearchPopResponseBodyData) *QueryQwenConferenceSgTicketSearchPopResponseBody {
	s.Data = v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBody) SetMessage(v string) *QueryQwenConferenceSgTicketSearchPopResponseBody {
	s.Message = &v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBody) SetRequestId(v string) *QueryQwenConferenceSgTicketSearchPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBody) SetSuccess(v bool) *QueryQwenConferenceSgTicketSearchPopResponseBody {
	s.Success = &v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryQwenConferenceSgTicketSearchPopResponseBodyData struct {
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
	// 12
	SubmitId *int64 `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
	// example:
	//
	// 1skladklasmda
	TicketToken *string `json:"TicketToken,omitempty" xml:"TicketToken,omitempty"`
}

func (s QueryQwenConferenceSgTicketSearchPopResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryQwenConferenceSgTicketSearchPopResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) GetCompanyName() *string {
	return s.CompanyName
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) GetExtFields() *string {
	return s.ExtFields
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) GetFirstName() *string {
	return s.FirstName
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) GetLastName() *string {
	return s.LastName
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) GetSubmitId() *int64 {
	return s.SubmitId
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) GetTicketToken() *string {
	return s.TicketToken
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) SetCompanyName(v string) *QueryQwenConferenceSgTicketSearchPopResponseBodyData {
	s.CompanyName = &v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) SetExtFields(v string) *QueryQwenConferenceSgTicketSearchPopResponseBodyData {
	s.ExtFields = &v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) SetFirstName(v string) *QueryQwenConferenceSgTicketSearchPopResponseBodyData {
	s.FirstName = &v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) SetLastName(v string) *QueryQwenConferenceSgTicketSearchPopResponseBodyData {
	s.LastName = &v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) SetSubmitId(v int64) *QueryQwenConferenceSgTicketSearchPopResponseBodyData {
	s.SubmitId = &v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) SetTicketToken(v string) *QueryQwenConferenceSgTicketSearchPopResponseBodyData {
	s.TicketToken = &v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponseBodyData) Validate() error {
	return dara.Validate(s)
}
