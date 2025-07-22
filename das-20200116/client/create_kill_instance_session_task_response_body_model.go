// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKillInstanceSessionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateKillInstanceSessionTaskResponseBody
	GetCode() *int64
	SetData(v string) *CreateKillInstanceSessionTaskResponseBody
	GetData() *string
	SetMessage(v string) *CreateKillInstanceSessionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateKillInstanceSessionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateKillInstanceSessionTaskResponseBody
	GetSuccess() *bool
}

type CreateKillInstanceSessionTaskResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the task that terminated the sessions.
	//
	// >  If the sessions of a PolarDB for MySQL cluster were terminated, **NodeId*	- is left empty, and **KillAllSessions*	- is set to **true**, the task IDs are returned based on the number of nodes. Example: ["f77d535b45405bd462b21caa3ee8\\*\\*\\*\\*", "e93ab549abb081eb5dcd5396a29b\\*\\*\\*\\*"].
	//
	// example:
	//
	// f77d535b45405bd462b21caa3ee8****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateKillInstanceSessionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKillInstanceSessionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKillInstanceSessionTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateKillInstanceSessionTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateKillInstanceSessionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKillInstanceSessionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKillInstanceSessionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateKillInstanceSessionTaskResponseBody) SetCode(v int64) *CreateKillInstanceSessionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKillInstanceSessionTaskResponseBody) SetData(v string) *CreateKillInstanceSessionTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateKillInstanceSessionTaskResponseBody) SetMessage(v string) *CreateKillInstanceSessionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKillInstanceSessionTaskResponseBody) SetRequestId(v string) *CreateKillInstanceSessionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKillInstanceSessionTaskResponseBody) SetSuccess(v bool) *CreateKillInstanceSessionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateKillInstanceSessionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
