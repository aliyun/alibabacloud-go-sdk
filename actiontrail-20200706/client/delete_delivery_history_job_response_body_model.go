// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeliveryHistoryJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDeliveryHistoryJobResponseBody
	GetRequestId() *string
}

type DeleteDeliveryHistoryJobResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D74DD20B-6598-429C-873B-B9B449B656B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeliveryHistoryJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeliveryHistoryJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryHistoryJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeliveryHistoryJobResponseBody) SetRequestId(v string) *DeleteDeliveryHistoryJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeliveryHistoryJobResponseBody) Validate() error {
	return dara.Validate(s)
}
