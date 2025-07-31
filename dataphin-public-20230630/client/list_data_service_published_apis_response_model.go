// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServicePublishedApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServicePublishedApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServicePublishedApisResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServicePublishedApisResponseBody) *ListDataServicePublishedApisResponse
	GetBody() *ListDataServicePublishedApisResponseBody
}

type ListDataServicePublishedApisResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServicePublishedApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServicePublishedApisResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponse) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServicePublishedApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServicePublishedApisResponse) GetBody() *ListDataServicePublishedApisResponseBody {
	return s.Body
}

func (s *ListDataServicePublishedApisResponse) SetHeaders(v map[string]*string) *ListDataServicePublishedApisResponse {
	s.Headers = v
	return s
}

func (s *ListDataServicePublishedApisResponse) SetStatusCode(v int32) *ListDataServicePublishedApisResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServicePublishedApisResponse) SetBody(v *ListDataServicePublishedApisResponseBody) *ListDataServicePublishedApisResponse {
	s.Body = v
	return s
}

func (s *ListDataServicePublishedApisResponse) Validate() error {
	return dara.Validate(s)
}
