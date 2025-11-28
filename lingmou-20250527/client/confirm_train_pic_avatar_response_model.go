// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmTrainPicAvatarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmTrainPicAvatarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmTrainPicAvatarResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmTrainPicAvatarResponseBody) *ConfirmTrainPicAvatarResponse
	GetBody() *ConfirmTrainPicAvatarResponseBody
}

type ConfirmTrainPicAvatarResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmTrainPicAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmTrainPicAvatarResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmTrainPicAvatarResponse) GoString() string {
	return s.String()
}

func (s *ConfirmTrainPicAvatarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmTrainPicAvatarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmTrainPicAvatarResponse) GetBody() *ConfirmTrainPicAvatarResponseBody {
	return s.Body
}

func (s *ConfirmTrainPicAvatarResponse) SetHeaders(v map[string]*string) *ConfirmTrainPicAvatarResponse {
	s.Headers = v
	return s
}

func (s *ConfirmTrainPicAvatarResponse) SetStatusCode(v int32) *ConfirmTrainPicAvatarResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmTrainPicAvatarResponse) SetBody(v *ConfirmTrainPicAvatarResponseBody) *ConfirmTrainPicAvatarResponse {
	s.Body = v
	return s
}

func (s *ConfirmTrainPicAvatarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
