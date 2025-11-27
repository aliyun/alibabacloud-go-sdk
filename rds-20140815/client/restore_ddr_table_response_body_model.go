// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreDdrTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *RestoreDdrTableResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *RestoreDdrTableResponseBody
	GetRequestId() *string
}

type RestoreDdrTableResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 75BBF1EF-1E90-4950-BFFB-252D26E8259B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreDdrTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreDdrTableResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreDdrTableResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RestoreDdrTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreDdrTableResponseBody) SetDBInstanceId(v string) *RestoreDdrTableResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *RestoreDdrTableResponseBody) SetRequestId(v string) *RestoreDdrTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreDdrTableResponseBody) Validate() error {
	return dara.Validate(s)
}
