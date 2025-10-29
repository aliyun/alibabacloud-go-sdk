// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNoTrainPicAvatarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNoTrainPicAvatarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNoTrainPicAvatarResponse
	GetStatusCode() *int32
	SetBody(v *CreateNoTrainPicAvatarResponseBody) *CreateNoTrainPicAvatarResponse
	GetBody() *CreateNoTrainPicAvatarResponseBody
}

type CreateNoTrainPicAvatarResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNoTrainPicAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNoTrainPicAvatarResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNoTrainPicAvatarResponse) GoString() string {
	return s.String()
}

func (s *CreateNoTrainPicAvatarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNoTrainPicAvatarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNoTrainPicAvatarResponse) GetBody() *CreateNoTrainPicAvatarResponseBody {
	return s.Body
}

func (s *CreateNoTrainPicAvatarResponse) SetHeaders(v map[string]*string) *CreateNoTrainPicAvatarResponse {
	s.Headers = v
	return s
}

func (s *CreateNoTrainPicAvatarResponse) SetStatusCode(v int32) *CreateNoTrainPicAvatarResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNoTrainPicAvatarResponse) SetBody(v *CreateNoTrainPicAvatarResponseBody) *CreateNoTrainPicAvatarResponse {
	s.Body = v
	return s
}

func (s *CreateNoTrainPicAvatarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
