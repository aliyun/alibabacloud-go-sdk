// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrainPicAvatarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrainPicAvatarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrainPicAvatarResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrainPicAvatarResponseBody) *CreateTrainPicAvatarResponse
	GetBody() *CreateTrainPicAvatarResponseBody
}

type CreateTrainPicAvatarResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrainPicAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrainPicAvatarResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainPicAvatarResponse) GoString() string {
	return s.String()
}

func (s *CreateTrainPicAvatarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrainPicAvatarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrainPicAvatarResponse) GetBody() *CreateTrainPicAvatarResponseBody {
	return s.Body
}

func (s *CreateTrainPicAvatarResponse) SetHeaders(v map[string]*string) *CreateTrainPicAvatarResponse {
	s.Headers = v
	return s
}

func (s *CreateTrainPicAvatarResponse) SetStatusCode(v int32) *CreateTrainPicAvatarResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrainPicAvatarResponse) SetBody(v *CreateTrainPicAvatarResponseBody) *CreateTrainPicAvatarResponse {
	s.Body = v
	return s
}

func (s *CreateTrainPicAvatarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
