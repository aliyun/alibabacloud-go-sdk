// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCmsExporterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *InstallCmsExporterResponseBody
	GetData() *string
	SetRequestId(v string) *InstallCmsExporterResponseBody
	GetRequestId() *string
}

type InstallCmsExporterResponseBody struct {
	// Indicates whether the call was successful.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// E7A04B0D-E2CA-59BB-8A9D-D5D349C22BF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallCmsExporterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallCmsExporterResponseBody) GoString() string {
	return s.String()
}

func (s *InstallCmsExporterResponseBody) GetData() *string {
	return s.Data
}

func (s *InstallCmsExporterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallCmsExporterResponseBody) SetData(v string) *InstallCmsExporterResponseBody {
	s.Data = &v
	return s
}

func (s *InstallCmsExporterResponseBody) SetRequestId(v string) *InstallCmsExporterResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallCmsExporterResponseBody) Validate() error {
	return dara.Validate(s)
}
