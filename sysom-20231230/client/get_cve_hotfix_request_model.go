// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCveHotfixRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCveIds(v []*string) *GetCveHotfixRequest
	GetCveIds() []*string
	SetKernelVersions(v []*string) *GetCveHotfixRequest
	GetKernelVersions() []*string
}

type GetCveHotfixRequest struct {
	// The list of CVE IDs.
	//
	// This parameter is required.
	CveIds []*string `json:"cveIds,omitempty" xml:"cveIds,omitempty" type:"Repeated"`
	// The list of full kernel versions.
	KernelVersions []*string `json:"kernelVersions,omitempty" xml:"kernelVersions,omitempty" type:"Repeated"`
}

func (s GetCveHotfixRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCveHotfixRequest) GoString() string {
	return s.String()
}

func (s *GetCveHotfixRequest) GetCveIds() []*string {
	return s.CveIds
}

func (s *GetCveHotfixRequest) GetKernelVersions() []*string {
	return s.KernelVersions
}

func (s *GetCveHotfixRequest) SetCveIds(v []*string) *GetCveHotfixRequest {
	s.CveIds = v
	return s
}

func (s *GetCveHotfixRequest) SetKernelVersions(v []*string) *GetCveHotfixRequest {
	s.KernelVersions = v
	return s
}

func (s *GetCveHotfixRequest) Validate() error {
	return dara.Validate(s)
}
