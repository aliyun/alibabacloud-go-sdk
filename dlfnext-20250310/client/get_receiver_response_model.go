// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReceiverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReceiverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReceiverResponse
	GetStatusCode() *int32
	SetBody(v *Receiver) *GetReceiverResponse
	GetBody() *Receiver
}

type GetReceiverResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Receiver          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReceiverResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReceiverResponse) GoString() string {
	return s.String()
}

func (s *GetReceiverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReceiverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReceiverResponse) GetBody() *Receiver {
	return s.Body
}

func (s *GetReceiverResponse) SetHeaders(v map[string]*string) *GetReceiverResponse {
	s.Headers = v
	return s
}

func (s *GetReceiverResponse) SetStatusCode(v int32) *GetReceiverResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReceiverResponse) SetBody(v *Receiver) *GetReceiverResponse {
	s.Body = v
	return s
}

func (s *GetReceiverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
