// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServicePublishedApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServicePublishedApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServicePublishedApiResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServicePublishedApiResponseBody) *GetDataServicePublishedApiResponse
	GetBody() *GetDataServicePublishedApiResponseBody
}

type GetDataServicePublishedApiResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServicePublishedApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServicePublishedApiResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServicePublishedApiResponse) GoString() string {
	return s.String()
}

func (s *GetDataServicePublishedApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServicePublishedApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServicePublishedApiResponse) GetBody() *GetDataServicePublishedApiResponseBody {
	return s.Body
}

func (s *GetDataServicePublishedApiResponse) SetHeaders(v map[string]*string) *GetDataServicePublishedApiResponse {
	s.Headers = v
	return s
}

func (s *GetDataServicePublishedApiResponse) SetStatusCode(v int32) *GetDataServicePublishedApiResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServicePublishedApiResponse) SetBody(v *GetDataServicePublishedApiResponseBody) *GetDataServicePublishedApiResponse {
	s.Body = v
	return s
}

func (s *GetDataServicePublishedApiResponse) Validate() error {
	return dara.Validate(s)
}
