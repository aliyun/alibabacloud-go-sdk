// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateChatSessionResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateChatSessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateChatSessionResponseBody
	GetRequestId() *string
	SetSession(v *UpdateChatSessionResponseBodySession) *UpdateChatSessionResponseBody
	GetSession() *UpdateChatSessionResponseBodySession
}

type UpdateChatSessionResponseBody struct {
	// The business status code. A value of 200 indicates success. A non-200 value indicates a backend error code (ERR.	- / InvalidParameter.*).
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description. This is empty when the request is successful.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The session ID.
	Session *UpdateChatSessionResponseBodySession `json:"session,omitempty" xml:"session,omitempty" type:"Struct"`
}

func (s UpdateChatSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatSessionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChatSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateChatSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateChatSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateChatSessionResponseBody) GetSession() *UpdateChatSessionResponseBodySession {
	return s.Session
}

func (s *UpdateChatSessionResponseBody) SetCode(v string) *UpdateChatSessionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateChatSessionResponseBody) SetMessage(v string) *UpdateChatSessionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateChatSessionResponseBody) SetRequestId(v string) *UpdateChatSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateChatSessionResponseBody) SetSession(v *UpdateChatSessionResponseBodySession) *UpdateChatSessionResponseBody {
	s.Session = v
	return s
}

func (s *UpdateChatSessionResponseBody) Validate() error {
	if s.Session != nil {
		if err := s.Session.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateChatSessionResponseBodySession struct {
	// The creation time.
	//
	// example:
	//
	// 1
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The message ID.
	//
	// example:
	//
	// exampleId
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the creation time exceeds 30 days.
	//
	// example:
	//
	// true
	IsExpired *bool `json:"isExpired,omitempty" xml:"isExpired,omitempty"`
	// The associated object ID.
	//
	// example:
	//
	// exampleObjectId
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The abstract model name used by the session (quick/standard/flagship).
	//
	// example:
	//
	// string_value
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The type.
	//
	// example:
	//
	// string_value
	Object *string `json:"object,omitempty" xml:"object,omitempty"`
	// The associated object ID.
	//
	// example:
	//
	// 2676
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The operating object name.
	//
	// example:
	//
	// string_value
	OperatingObjectName []*string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty" type:"Repeated"`
	// The title.
	//
	// example:
	//
	// Sample title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s UpdateChatSessionResponseBodySession) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatSessionResponseBodySession) GoString() string {
	return s.String()
}

func (s *UpdateChatSessionResponseBodySession) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *UpdateChatSessionResponseBodySession) GetId() *string {
	return s.Id
}

func (s *UpdateChatSessionResponseBodySession) GetIsExpired() *bool {
	return s.IsExpired
}

func (s *UpdateChatSessionResponseBodySession) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *UpdateChatSessionResponseBodySession) GetModel() *string {
	return s.Model
}

func (s *UpdateChatSessionResponseBodySession) GetObject() *string {
	return s.Object
}

func (s *UpdateChatSessionResponseBodySession) GetObjectId() *string {
	return s.ObjectId
}

func (s *UpdateChatSessionResponseBodySession) GetOperatingObjectName() []*string {
	return s.OperatingObjectName
}

func (s *UpdateChatSessionResponseBodySession) GetTitle() *string {
	return s.Title
}

func (s *UpdateChatSessionResponseBodySession) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *UpdateChatSessionResponseBodySession) SetCreatedAt(v int64) *UpdateChatSessionResponseBodySession {
	s.CreatedAt = &v
	return s
}

func (s *UpdateChatSessionResponseBodySession) SetId(v string) *UpdateChatSessionResponseBodySession {
	s.Id = &v
	return s
}

func (s *UpdateChatSessionResponseBodySession) SetIsExpired(v bool) *UpdateChatSessionResponseBodySession {
	s.IsExpired = &v
	return s
}

func (s *UpdateChatSessionResponseBodySession) SetMetadata(v map[string]interface{}) *UpdateChatSessionResponseBodySession {
	s.Metadata = v
	return s
}

func (s *UpdateChatSessionResponseBodySession) SetModel(v string) *UpdateChatSessionResponseBodySession {
	s.Model = &v
	return s
}

func (s *UpdateChatSessionResponseBodySession) SetObject(v string) *UpdateChatSessionResponseBodySession {
	s.Object = &v
	return s
}

func (s *UpdateChatSessionResponseBodySession) SetObjectId(v string) *UpdateChatSessionResponseBodySession {
	s.ObjectId = &v
	return s
}

func (s *UpdateChatSessionResponseBodySession) SetOperatingObjectName(v []*string) *UpdateChatSessionResponseBodySession {
	s.OperatingObjectName = v
	return s
}

func (s *UpdateChatSessionResponseBodySession) SetTitle(v string) *UpdateChatSessionResponseBodySession {
	s.Title = &v
	return s
}

func (s *UpdateChatSessionResponseBodySession) SetUpdatedAt(v int64) *UpdateChatSessionResponseBodySession {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateChatSessionResponseBodySession) Validate() error {
	return dara.Validate(s)
}
