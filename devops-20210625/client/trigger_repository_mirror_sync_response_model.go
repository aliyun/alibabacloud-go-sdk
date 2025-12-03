// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerRepositoryMirrorSyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerRepositoryMirrorSyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerRepositoryMirrorSyncResponse
	GetStatusCode() *int32
	SetBody(v *TriggerRepositoryMirrorSyncResponseBody) *TriggerRepositoryMirrorSyncResponse
	GetBody() *TriggerRepositoryMirrorSyncResponseBody
}

type TriggerRepositoryMirrorSyncResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerRepositoryMirrorSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerRepositoryMirrorSyncResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerRepositoryMirrorSyncResponse) GoString() string {
	return s.String()
}

func (s *TriggerRepositoryMirrorSyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerRepositoryMirrorSyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerRepositoryMirrorSyncResponse) GetBody() *TriggerRepositoryMirrorSyncResponseBody {
	return s.Body
}

func (s *TriggerRepositoryMirrorSyncResponse) SetHeaders(v map[string]*string) *TriggerRepositoryMirrorSyncResponse {
	s.Headers = v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponse) SetStatusCode(v int32) *TriggerRepositoryMirrorSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponse) SetBody(v *TriggerRepositoryMirrorSyncResponseBody) *TriggerRepositoryMirrorSyncResponse {
	s.Body = v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
