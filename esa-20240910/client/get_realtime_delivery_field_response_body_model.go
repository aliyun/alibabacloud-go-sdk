// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeDeliveryFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFieldContent(v map[string]*FieldContentValue) *GetRealtimeDeliveryFieldResponseBody
	GetFieldContent() map[string]*FieldContentValue
	SetRequestId(v string) *GetRealtimeDeliveryFieldResponseBody
	GetRequestId() *string
}

type GetRealtimeDeliveryFieldResponseBody struct {
	// The fields returned.
	FieldContent map[string]*FieldContentValue `json:"FieldContent,omitempty" xml:"FieldContent,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRealtimeDeliveryFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeDeliveryFieldResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealtimeDeliveryFieldResponseBody) GetFieldContent() map[string]*FieldContentValue {
	return s.FieldContent
}

func (s *GetRealtimeDeliveryFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRealtimeDeliveryFieldResponseBody) SetFieldContent(v map[string]*FieldContentValue) *GetRealtimeDeliveryFieldResponseBody {
	s.FieldContent = v
	return s
}

func (s *GetRealtimeDeliveryFieldResponseBody) SetRequestId(v string) *GetRealtimeDeliveryFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealtimeDeliveryFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
