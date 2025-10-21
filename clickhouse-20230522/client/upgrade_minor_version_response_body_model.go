// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMinorVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpgradeMinorVersionResponseBodyData) *UpgradeMinorVersionResponseBody
	GetData() *UpgradeMinorVersionResponseBodyData
	SetRequestId(v string) *UpgradeMinorVersionResponseBody
	GetRequestId() *string
}

type UpgradeMinorVersionResponseBody struct {
	// The returned result.
	Data *UpgradeMinorVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FE242962-6DA3-5FC8-9691-37B62A3210F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeMinorVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponseBody) GetData() *UpgradeMinorVersionResponseBodyData {
	return s.Data
}

func (s *UpgradeMinorVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeMinorVersionResponseBody) SetData(v *UpgradeMinorVersionResponseBodyData) *UpgradeMinorVersionResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeMinorVersionResponseBody) SetRequestId(v string) *UpgradeMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeMinorVersionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeMinorVersionResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// cc-uf6x229yeq166****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s UpgradeMinorVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMinorVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpgradeMinorVersionResponseBodyData) SetDBInstanceName(v string) *UpgradeMinorVersionResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeMinorVersionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
