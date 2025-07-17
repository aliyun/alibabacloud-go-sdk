// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeLhDagOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeLhDagOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeLhDagOwnerResponse
	GetStatusCode() *int32
	SetBody(v *ChangeLhDagOwnerResponseBody) *ChangeLhDagOwnerResponse
	GetBody() *ChangeLhDagOwnerResponseBody
}

type ChangeLhDagOwnerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeLhDagOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeLhDagOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeLhDagOwnerResponse) GoString() string {
	return s.String()
}

func (s *ChangeLhDagOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeLhDagOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeLhDagOwnerResponse) GetBody() *ChangeLhDagOwnerResponseBody {
	return s.Body
}

func (s *ChangeLhDagOwnerResponse) SetHeaders(v map[string]*string) *ChangeLhDagOwnerResponse {
	s.Headers = v
	return s
}

func (s *ChangeLhDagOwnerResponse) SetStatusCode(v int32) *ChangeLhDagOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeLhDagOwnerResponse) SetBody(v *ChangeLhDagOwnerResponseBody) *ChangeLhDagOwnerResponse {
	s.Body = v
	return s
}

func (s *ChangeLhDagOwnerResponse) Validate() error {
	return dara.Validate(s)
}
