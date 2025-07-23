// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigBackupTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigBackupTaskResponseBody
	GetRequestId() *string
}

type ConfigBackupTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigBackupTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigBackupTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigBackupTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigBackupTaskResponseBody) SetRequestId(v string) *ConfigBackupTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigBackupTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
