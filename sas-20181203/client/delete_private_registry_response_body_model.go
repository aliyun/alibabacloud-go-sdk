// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateRegistryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeletePrivateRegistryResponseBody
	GetData() *bool
	SetRequestId(v string) *DeletePrivateRegistryResponseBody
	GetRequestId() *string
}

type DeletePrivateRegistryResponseBody struct {
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 818E3B38-F018-50FF-9A85-5A521747****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrivateRegistryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateRegistryResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivateRegistryResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeletePrivateRegistryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrivateRegistryResponseBody) SetData(v bool) *DeletePrivateRegistryResponseBody {
	s.Data = &v
	return s
}

func (s *DeletePrivateRegistryResponseBody) SetRequestId(v string) *DeletePrivateRegistryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrivateRegistryResponseBody) Validate() error {
	return dara.Validate(s)
}
