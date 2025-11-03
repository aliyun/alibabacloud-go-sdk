// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCopilotHistoryResponseBody
	GetCode() *string
	SetData(v []*GetCopilotHistoryResponseBodyData) *GetCopilotHistoryResponseBody
	GetData() []*GetCopilotHistoryResponseBodyData
	SetMessage(v string) *GetCopilotHistoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCopilotHistoryResponseBody
	GetRequestId() *string
}

type GetCopilotHistoryResponseBody struct {
	// example:
	//
	// SysomOpenAPI.InvalidParameter
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetCopilotHistoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCopilotHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetCopilotHistoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCopilotHistoryResponseBody) GetData() []*GetCopilotHistoryResponseBodyData {
	return s.Data
}

func (s *GetCopilotHistoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCopilotHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCopilotHistoryResponseBody) SetCode(v string) *GetCopilotHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetCopilotHistoryResponseBody) SetData(v []*GetCopilotHistoryResponseBodyData) *GetCopilotHistoryResponseBody {
	s.Data = v
	return s
}

func (s *GetCopilotHistoryResponseBody) SetMessage(v string) *GetCopilotHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetCopilotHistoryResponseBody) SetRequestId(v string) *GetCopilotHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCopilotHistoryResponseBody) Validate() error {
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

type GetCopilotHistoryResponseBodyData struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2024-09-02 10:02:39
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// user
	//
	// copilot
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s GetCopilotHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCopilotHistoryResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetCopilotHistoryResponseBodyData) GetTime() *string {
	return s.Time
}

func (s *GetCopilotHistoryResponseBodyData) GetUser() *string {
	return s.User
}

func (s *GetCopilotHistoryResponseBodyData) SetContent(v string) *GetCopilotHistoryResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetCopilotHistoryResponseBodyData) SetTime(v string) *GetCopilotHistoryResponseBodyData {
	s.Time = &v
	return s
}

func (s *GetCopilotHistoryResponseBodyData) SetUser(v string) *GetCopilotHistoryResponseBodyData {
	s.User = &v
	return s
}

func (s *GetCopilotHistoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
