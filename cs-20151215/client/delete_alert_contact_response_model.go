// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertContactResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlertContactResponseBody) *DeleteAlertContactResponse
	GetBody() *DeleteAlertContactResponseBody
}

type DeleteAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s DeleteAlertContactResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertContactResponse) GetBody() *DeleteAlertContactResponseBody {
	return s.Body
}

func (s *DeleteAlertContactResponse) SetHeaders(v map[string]*string) *DeleteAlertContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertContactResponse) SetStatusCode(v int32) *DeleteAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertContactResponse) SetBody(v *DeleteAlertContactResponseBody) *DeleteAlertContactResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertContactResponse) Validate() error {
	return dara.Validate(s)
}

type DeleteAlertContactResponseBody struct {
	Result []*DeleteAlertContactResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s DeleteAlertContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactResponseBody) GetResult() []*DeleteAlertContactResponseBodyResult {
	return s.Result
}

func (s *DeleteAlertContactResponseBody) SetResult(v []*DeleteAlertContactResponseBodyResult) *DeleteAlertContactResponseBody {
	s.Result = v
	return s
}

func (s *DeleteAlertContactResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteAlertContactResponseBodyResult struct {
	// The deletion status.
	//
	// 	- true: The alert contact was deleted.
	//
	// 	- false: The alert contact failed to be deleted.
	//
	// example:
	//
	// true
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// Delete contact resource failed.
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// An alert contact ID.
	//
	// example:
	//
	// 12345
	ContactId *string `json:"contact_id,omitempty" xml:"contact_id,omitempty"`
}

func (s DeleteAlertContactResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactResponseBodyResult) GetStatus() *bool {
	return s.Status
}

func (s *DeleteAlertContactResponseBodyResult) GetMsg() *string {
	return s.Msg
}

func (s *DeleteAlertContactResponseBodyResult) GetContactId() *string {
	return s.ContactId
}

func (s *DeleteAlertContactResponseBodyResult) SetStatus(v bool) *DeleteAlertContactResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DeleteAlertContactResponseBodyResult) SetMsg(v string) *DeleteAlertContactResponseBodyResult {
	s.Msg = &v
	return s
}

func (s *DeleteAlertContactResponseBodyResult) SetContactId(v string) *DeleteAlertContactResponseBodyResult {
	s.ContactId = &v
	return s
}

func (s *DeleteAlertContactResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
