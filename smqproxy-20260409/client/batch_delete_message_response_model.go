// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteMessageResponse
	GetStatusCode() *int32
}

type BatchDeleteMessageResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s BatchDeleteMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteMessageResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteMessageResponse) SetHeaders(v map[string]*string) *BatchDeleteMessageResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteMessageResponse) SetStatusCode(v int32) *BatchDeleteMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteMessageResponse) Validate() error {
	return dara.Validate(s)
}
