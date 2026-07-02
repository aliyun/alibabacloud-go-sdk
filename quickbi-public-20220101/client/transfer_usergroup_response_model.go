// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferUsergroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferUsergroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferUsergroupResponse
	GetStatusCode() *int32
	SetBody(v *TransferUsergroupResponseBody) *TransferUsergroupResponse
	GetBody() *TransferUsergroupResponseBody
}

type TransferUsergroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferUsergroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferUsergroupResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferUsergroupResponse) GoString() string {
	return s.String()
}

func (s *TransferUsergroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferUsergroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferUsergroupResponse) GetBody() *TransferUsergroupResponseBody {
	return s.Body
}

func (s *TransferUsergroupResponse) SetHeaders(v map[string]*string) *TransferUsergroupResponse {
	s.Headers = v
	return s
}

func (s *TransferUsergroupResponse) SetStatusCode(v int32) *TransferUsergroupResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferUsergroupResponse) SetBody(v *TransferUsergroupResponseBody) *TransferUsergroupResponse {
	s.Body = v
	return s
}

func (s *TransferUsergroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
