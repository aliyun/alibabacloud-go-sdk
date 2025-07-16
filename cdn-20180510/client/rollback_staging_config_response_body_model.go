// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackStagingConfigResponseBody
	GetRequestId() *string
}

type RollbackStagingConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackStagingConfigResponseBody) SetRequestId(v string) *RollbackStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackStagingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
