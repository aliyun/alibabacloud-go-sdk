// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcid(v string) *GetRecordUrlRequest
	GetAcid() *string
	SetInstanceId(v string) *GetRecordUrlRequest
	GetInstanceId() *string
	SetRecordType(v string) *GetRecordUrlRequest
	GetRecordType() *string
}

type GetRecordUrlRequest struct {
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001067****
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID. You can obtain it in the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview) > Instance Management.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The call type. Valid values:
	//
	// - **DUP_CALL**: Incoming and outgoing calls.
	//
	// - **IVR_CALL**: IVR outbound call.
	//
	// - **SMART_CALL**: Intelligent outbound call.
	//
	// This parameter is required.
	//
	// example:
	//
	// DUP_CALL
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
}

func (s GetRecordUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecordUrlRequest) GoString() string {
	return s.String()
}

func (s *GetRecordUrlRequest) GetAcid() *string {
	return s.Acid
}

func (s *GetRecordUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRecordUrlRequest) GetRecordType() *string {
	return s.RecordType
}

func (s *GetRecordUrlRequest) SetAcid(v string) *GetRecordUrlRequest {
	s.Acid = &v
	return s
}

func (s *GetRecordUrlRequest) SetInstanceId(v string) *GetRecordUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRecordUrlRequest) SetRecordType(v string) *GetRecordUrlRequest {
	s.RecordType = &v
	return s
}

func (s *GetRecordUrlRequest) Validate() error {
	return dara.Validate(s)
}
