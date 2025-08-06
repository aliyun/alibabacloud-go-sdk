// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackVodStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackVodStagingConfigResponseBody
	GetRequestId() *string
}

type RollbackVodStagingConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackVodStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackVodStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackVodStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackVodStagingConfigResponseBody) SetRequestId(v string) *RollbackVodStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackVodStagingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
