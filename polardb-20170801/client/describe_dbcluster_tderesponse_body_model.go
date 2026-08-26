// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterTDEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutomaticRotation(v string) *DescribeDBClusterTDEResponseBody
	GetAutomaticRotation() *string
	SetDBClusterId(v string) *DescribeDBClusterTDEResponseBody
	GetDBClusterId() *string
	SetEncryptNewTables(v string) *DescribeDBClusterTDEResponseBody
	GetEncryptNewTables() *string
	SetEncryptionKey(v string) *DescribeDBClusterTDEResponseBody
	GetEncryptionKey() *string
	SetEncryptionKeyStatus(v string) *DescribeDBClusterTDEResponseBody
	GetEncryptionKeyStatus() *string
	SetRequestId(v string) *DescribeDBClusterTDEResponseBody
	GetRequestId() *string
	SetRotationInterval(v string) *DescribeDBClusterTDEResponseBody
	GetRotationInterval() *string
	SetTDERegion(v string) *DescribeDBClusterTDEResponseBody
	GetTDERegion() *string
	SetTDEStatus(v string) *DescribeDBClusterTDEResponseBody
	GetTDEStatus() *string
}

type DescribeDBClusterTDEResponseBody struct {
	// Indicates whether automatic key rotation is enabled. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// > This parameter is returned only when the database engine is PostgreSQL-compatible or Oracle-syntax-compatible.
	//
	// example:
	//
	// Enabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether automatic encryption of all newly created tables is enabled. Valid values:
	//
	// - **ON**: Enabled.
	//
	// - **OFF**: Disabled.
	//
	// > This parameter is returned only when the database engine is MySQL-compatible.
	//
	// example:
	//
	// ON
	EncryptNewTables *string `json:"EncryptNewTables,omitempty" xml:"EncryptNewTables,omitempty"`
	// The custom key ID.
	//
	// example:
	//
	// 2a4f4ac2-****-****-****-************
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The status of the key. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	EncryptionKeyStatus *string `json:"EncryptionKeyStatus,omitempty" xml:"EncryptionKeyStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E37D1508-EC3B-4E06-A24A-C7AC31******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The automatic key rotation interval configured in Key Management Service (KMS). If no automatic key rotation interval is configured, 0 s is returned. Unit: s.
	//
	// For example, if the rotation interval is 7 days, 604800 s is returned.
	//
	// > This parameter is returned only when the database engine is PostgreSQL-compatible or Oracle-syntax-compatible and the value of AutomaticRotation is Enabled.
	//
	// example:
	//
	// 604800s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The region where the TDE key resides.
	//
	// example:
	//
	// cn-beijing
	TDERegion *string `json:"TDERegion,omitempty" xml:"TDERegion,omitempty"`
	// Indicates whether TDE is enabled. Valid values:
	//
	// 	- **Enabled**: Enabled.
	//
	// 	- **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	TDEStatus *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s DescribeDBClusterTDEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterTDEResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterTDEResponseBody) GetAutomaticRotation() *string {
	return s.AutomaticRotation
}

func (s *DescribeDBClusterTDEResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterTDEResponseBody) GetEncryptNewTables() *string {
	return s.EncryptNewTables
}

func (s *DescribeDBClusterTDEResponseBody) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeDBClusterTDEResponseBody) GetEncryptionKeyStatus() *string {
	return s.EncryptionKeyStatus
}

func (s *DescribeDBClusterTDEResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterTDEResponseBody) GetRotationInterval() *string {
	return s.RotationInterval
}

func (s *DescribeDBClusterTDEResponseBody) GetTDERegion() *string {
	return s.TDERegion
}

func (s *DescribeDBClusterTDEResponseBody) GetTDEStatus() *string {
	return s.TDEStatus
}

func (s *DescribeDBClusterTDEResponseBody) SetAutomaticRotation(v string) *DescribeDBClusterTDEResponseBody {
	s.AutomaticRotation = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) SetDBClusterId(v string) *DescribeDBClusterTDEResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) SetEncryptNewTables(v string) *DescribeDBClusterTDEResponseBody {
	s.EncryptNewTables = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) SetEncryptionKey(v string) *DescribeDBClusterTDEResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) SetEncryptionKeyStatus(v string) *DescribeDBClusterTDEResponseBody {
	s.EncryptionKeyStatus = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) SetRequestId(v string) *DescribeDBClusterTDEResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) SetRotationInterval(v string) *DescribeDBClusterTDEResponseBody {
	s.RotationInterval = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) SetTDERegion(v string) *DescribeDBClusterTDEResponseBody {
	s.TDERegion = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) SetTDEStatus(v string) *DescribeDBClusterTDEResponseBody {
	s.TDEStatus = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) Validate() error {
	return dara.Validate(s)
}
