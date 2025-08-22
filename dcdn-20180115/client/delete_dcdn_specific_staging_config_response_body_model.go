// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnSpecificStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnSpecificStagingConfigResponseBody
	GetRequestId() *string
}

type DeleteDcdnSpecificStagingConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnSpecificStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnSpecificStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSpecificStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnSpecificStagingConfigResponseBody) SetRequestId(v string) *DeleteDcdnSpecificStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnSpecificStagingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
