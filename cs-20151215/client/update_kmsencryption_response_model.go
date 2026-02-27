// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKMSEncryptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKMSEncryptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKMSEncryptionResponse
	GetStatusCode() *int32
}

type UpdateKMSEncryptionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateKMSEncryptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKMSEncryptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateKMSEncryptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKMSEncryptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKMSEncryptionResponse) SetHeaders(v map[string]*string) *UpdateKMSEncryptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateKMSEncryptionResponse) SetStatusCode(v int32) *UpdateKMSEncryptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKMSEncryptionResponse) Validate() error {
	return dara.Validate(s)
}
