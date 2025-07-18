// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtensions(v string) *CreateExtensionsResponseBody
	GetExtensions() *string
	SetRequestId(v string) *CreateExtensionsResponseBody
	GetRequestId() *string
}

type CreateExtensionsResponseBody struct {
	// The name of the extension that you want to install. Multiple extension names are separated with commas (,).
	//
	// example:
	//
	// citext, dblink
	Extensions *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExtensionsResponseBody) GetExtensions() *string {
	return s.Extensions
}

func (s *CreateExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExtensionsResponseBody) SetExtensions(v string) *CreateExtensionsResponseBody {
	s.Extensions = &v
	return s
}

func (s *CreateExtensionsResponseBody) SetRequestId(v string) *CreateExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExtensionsResponseBody) Validate() error {
	return dara.Validate(s)
}
