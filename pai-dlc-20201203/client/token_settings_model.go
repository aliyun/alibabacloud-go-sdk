// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTokenSettings interface {
	dara.Model
	String() string
	GoString() string
	SetEnableCrossAccountAccess(v bool) *TokenSettings
	GetEnableCrossAccountAccess() *bool
	SetEnableLogDownloadJob(v bool) *TokenSettings
	GetEnableLogDownloadJob() *bool
}

type TokenSettings struct {
	EnableCrossAccountAccess *bool `json:"EnableCrossAccountAccess,omitempty" xml:"EnableCrossAccountAccess,omitempty"`
	EnableLogDownloadJob     *bool `json:"EnableLogDownloadJob,omitempty" xml:"EnableLogDownloadJob,omitempty"`
}

func (s TokenSettings) String() string {
	return dara.Prettify(s)
}

func (s TokenSettings) GoString() string {
	return s.String()
}

func (s *TokenSettings) GetEnableCrossAccountAccess() *bool {
	return s.EnableCrossAccountAccess
}

func (s *TokenSettings) GetEnableLogDownloadJob() *bool {
	return s.EnableLogDownloadJob
}

func (s *TokenSettings) SetEnableCrossAccountAccess(v bool) *TokenSettings {
	s.EnableCrossAccountAccess = &v
	return s
}

func (s *TokenSettings) SetEnableLogDownloadJob(v bool) *TokenSettings {
	s.EnableLogDownloadJob = &v
	return s
}

func (s *TokenSettings) Validate() error {
	return dara.Validate(s)
}
