// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackDcdnStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackDcdnStagingConfigResponseBody
	GetRequestId() *string
}

type RollbackDcdnStagingConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackDcdnStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackDcdnStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackDcdnStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackDcdnStagingConfigResponseBody) SetRequestId(v string) *RollbackDcdnStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackDcdnStagingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
