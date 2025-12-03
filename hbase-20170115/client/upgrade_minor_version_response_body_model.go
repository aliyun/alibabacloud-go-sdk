// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMinorVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNewVersion(v string) *UpgradeMinorVersionResponseBody
	GetNewVersion() *string
	SetOldVersion(v string) *UpgradeMinorVersionResponseBody
	GetOldVersion() *string
	SetRequestId(v string) *UpgradeMinorVersionResponseBody
	GetRequestId() *string
}

type UpgradeMinorVersionResponseBody struct {
	NewVersion *string `json:"NewVersion,omitempty" xml:"NewVersion,omitempty"`
	OldVersion *string `json:"OldVersion,omitempty" xml:"OldVersion,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeMinorVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponseBody) GetNewVersion() *string {
	return s.NewVersion
}

func (s *UpgradeMinorVersionResponseBody) GetOldVersion() *string {
	return s.OldVersion
}

func (s *UpgradeMinorVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeMinorVersionResponseBody) SetNewVersion(v string) *UpgradeMinorVersionResponseBody {
	s.NewVersion = &v
	return s
}

func (s *UpgradeMinorVersionResponseBody) SetOldVersion(v string) *UpgradeMinorVersionResponseBody {
	s.OldVersion = &v
	return s
}

func (s *UpgradeMinorVersionResponseBody) SetRequestId(v string) *UpgradeMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeMinorVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
