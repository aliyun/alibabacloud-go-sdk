// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyInfoInRecycleBinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessKeyInfoInRecycleBinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessKeyInfoInRecycleBinResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessKeyInfoInRecycleBinResponseBody) *GetAccessKeyInfoInRecycleBinResponse
	GetBody() *GetAccessKeyInfoInRecycleBinResponseBody
}

type GetAccessKeyInfoInRecycleBinResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessKeyInfoInRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessKeyInfoInRecycleBinResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyInfoInRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyInfoInRecycleBinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessKeyInfoInRecycleBinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessKeyInfoInRecycleBinResponse) GetBody() *GetAccessKeyInfoInRecycleBinResponseBody {
	return s.Body
}

func (s *GetAccessKeyInfoInRecycleBinResponse) SetHeaders(v map[string]*string) *GetAccessKeyInfoInRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponse) SetStatusCode(v int32) *GetAccessKeyInfoInRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponse) SetBody(v *GetAccessKeyInfoInRecycleBinResponseBody) *GetAccessKeyInfoInRecycleBinResponse {
	s.Body = v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
