// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLastUpgradeRecordRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetLastUpgradeRecordRequest struct {
}

func (s GetLastUpgradeRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLastUpgradeRecordRequest) GoString() string {
	return s.String()
}

func (s *GetLastUpgradeRecordRequest) Validate() error {
	return dara.Validate(s)
}
