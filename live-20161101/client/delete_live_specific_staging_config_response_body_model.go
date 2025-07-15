// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSpecificStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveSpecificStagingConfigResponseBody
	GetRequestId() *string
}

type DeleteLiveSpecificStagingConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveSpecificStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSpecificStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveSpecificStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveSpecificStagingConfigResponseBody) SetRequestId(v string) *DeleteLiveSpecificStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveSpecificStagingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
