// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrainPicAvatarStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrainPicAvatarStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrainPicAvatarStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetTrainPicAvatarStatusResponseBody) *GetTrainPicAvatarStatusResponse
	GetBody() *GetTrainPicAvatarStatusResponseBody
}

type GetTrainPicAvatarStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrainPicAvatarStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrainPicAvatarStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrainPicAvatarStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTrainPicAvatarStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrainPicAvatarStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrainPicAvatarStatusResponse) GetBody() *GetTrainPicAvatarStatusResponseBody {
	return s.Body
}

func (s *GetTrainPicAvatarStatusResponse) SetHeaders(v map[string]*string) *GetTrainPicAvatarStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTrainPicAvatarStatusResponse) SetStatusCode(v int32) *GetTrainPicAvatarStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrainPicAvatarStatusResponse) SetBody(v *GetTrainPicAvatarStatusResponseBody) *GetTrainPicAvatarStatusResponse {
	s.Body = v
	return s
}

func (s *GetTrainPicAvatarStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
