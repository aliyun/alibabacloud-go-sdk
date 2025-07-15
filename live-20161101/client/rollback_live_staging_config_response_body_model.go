// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackLiveStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackLiveStagingConfigResponseBody
	GetRequestId() *string
}

type RollbackLiveStagingConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackLiveStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackLiveStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackLiveStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackLiveStagingConfigResponseBody) SetRequestId(v string) *RollbackLiveStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackLiveStagingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
