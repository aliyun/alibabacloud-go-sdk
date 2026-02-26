// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceReplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDBInstanceReplicationResponseBody
	GetRequestId() *string
}

type CreateDBInstanceReplicationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 76364A52-E0AB-5CC8-9818-CF1DC482C092
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceReplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceReplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceReplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstanceReplicationResponseBody) SetRequestId(v string) *CreateDBInstanceReplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceReplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
