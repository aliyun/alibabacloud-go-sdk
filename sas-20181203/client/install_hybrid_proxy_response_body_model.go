// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallHybridProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallHybridProxyResponseBody
	GetRequestId() *string
}

type InstallHybridProxyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E55BE5DB-E2DF-57EB-A0C5-7A85EEA67A4C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallHybridProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallHybridProxyResponseBody) GoString() string {
	return s.String()
}

func (s *InstallHybridProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallHybridProxyResponseBody) SetRequestId(v string) *InstallHybridProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallHybridProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
