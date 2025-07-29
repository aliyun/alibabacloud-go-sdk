// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertContactGroupResponse
	GetStatusCode() *int32
	SetBody(v []*DeleteAlertContactGroupResponseBody) *DeleteAlertContactGroupResponse
	GetBody() []*DeleteAlertContactGroupResponseBody
}

type DeleteAlertContactGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*DeleteAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s DeleteAlertContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertContactGroupResponse) GetBody() []*DeleteAlertContactGroupResponseBody {
	return s.Body
}

func (s *DeleteAlertContactGroupResponse) SetHeaders(v map[string]*string) *DeleteAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertContactGroupResponse) SetStatusCode(v int32) *DeleteAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertContactGroupResponse) SetBody(v []*DeleteAlertContactGroupResponseBody) *DeleteAlertContactGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertContactGroupResponse) Validate() error {
	return dara.Validate(s)
}

type DeleteAlertContactGroupResponseBody struct {
	// The deletion status.
	//
	// 	- true: The alert contact group was deleted.
	//
	// 	- false: The alert contact group failed to be deleted.
	//
	// example:
	//
	// true
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// Delete contact group resource failed.
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The alert contact group ID.
	//
	// example:
	//
	// 12345
	ContactGroupId *string `json:"contact_group_id,omitempty" xml:"contact_group_id,omitempty"`
}

func (s DeleteAlertContactGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *DeleteAlertContactGroupResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DeleteAlertContactGroupResponseBody) GetContactGroupId() *string {
	return s.ContactGroupId
}

func (s *DeleteAlertContactGroupResponseBody) SetStatus(v bool) *DeleteAlertContactGroupResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteAlertContactGroupResponseBody) SetMsg(v string) *DeleteAlertContactGroupResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteAlertContactGroupResponseBody) SetContactGroupId(v string) *DeleteAlertContactGroupResponseBody {
	s.ContactGroupId = &v
	return s
}

func (s *DeleteAlertContactGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
