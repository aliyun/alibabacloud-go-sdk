// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEntityStoreDataHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetEntityStoreDataHeaders
	GetCommonHeaders() map[string]*string
	SetAcceptEncoding(v string) *GetEntityStoreDataHeaders
	GetAcceptEncoding() *string
}

type GetEntityStoreDataHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// gzip
	AcceptEncoding *string `json:"acceptEncoding,omitempty" xml:"acceptEncoding,omitempty"`
}

func (s GetEntityStoreDataHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetEntityStoreDataHeaders) GoString() string {
	return s.String()
}

func (s *GetEntityStoreDataHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetEntityStoreDataHeaders) GetAcceptEncoding() *string {
	return s.AcceptEncoding
}

func (s *GetEntityStoreDataHeaders) SetCommonHeaders(v map[string]*string) *GetEntityStoreDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEntityStoreDataHeaders) SetAcceptEncoding(v string) *GetEntityStoreDataHeaders {
	s.AcceptEncoding = &v
	return s
}

func (s *GetEntityStoreDataHeaders) Validate() error {
	return dara.Validate(s)
}
