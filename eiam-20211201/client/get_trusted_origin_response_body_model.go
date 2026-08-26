// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrustedOriginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTrustedOriginResponseBody
	GetRequestId() *string
	SetTrustedOrigin(v *GetTrustedOriginResponseBodyTrustedOrigin) *GetTrustedOriginResponseBody
	GetTrustedOrigin() *GetTrustedOriginResponseBodyTrustedOrigin
}

type GetTrustedOriginResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-example
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The trusted origin.
	TrustedOrigin *GetTrustedOriginResponseBodyTrustedOrigin `json:"TrustedOrigin,omitempty" xml:"TrustedOrigin,omitempty" type:"Struct"`
}

func (s GetTrustedOriginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrustedOriginResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrustedOriginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrustedOriginResponseBody) GetTrustedOrigin() *GetTrustedOriginResponseBodyTrustedOrigin {
	return s.TrustedOrigin
}

func (s *GetTrustedOriginResponseBody) SetRequestId(v string) *GetTrustedOriginResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrustedOriginResponseBody) SetTrustedOrigin(v *GetTrustedOriginResponseBodyTrustedOrigin) *GetTrustedOriginResponseBody {
	s.TrustedOrigin = v
	return s
}

func (s *GetTrustedOriginResponseBody) Validate() error {
	if s.TrustedOrigin != nil {
		if err := s.TrustedOrigin.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTrustedOriginResponseBodyTrustedOrigin struct {
	// The creation time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-08-20T08:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_example
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The browser origin.
	//
	// example:
	//
	// https://console.qoder.com
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The status.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the trusted origin.
	//
	// example:
	//
	// Qoder Production Console
	TrustOriginName *string `json:"TrustOriginName,omitempty" xml:"TrustOriginName,omitempty"`
	// The ID of the trusted origin.
	//
	// example:
	//
	// to_example
	TrustedOriginId *string `json:"TrustedOriginId,omitempty" xml:"TrustedOriginId,omitempty"`
	// The trusted origin scene.
	TrustedOriginScene []*string `json:"TrustedOriginScene,omitempty" xml:"TrustedOriginScene,omitempty" type:"Repeated"`
	// The update time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-08-20T08:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetTrustedOriginResponseBodyTrustedOrigin) String() string {
	return dara.Prettify(s)
}

func (s GetTrustedOriginResponseBodyTrustedOrigin) GoString() string {
	return s.String()
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) GetOrigin() *string {
	return s.Origin
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) GetStatus() *string {
	return s.Status
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) GetTrustOriginName() *string {
	return s.TrustOriginName
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) GetTrustedOriginId() *string {
	return s.TrustedOriginId
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) GetTrustedOriginScene() []*string {
	return s.TrustedOriginScene
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) SetCreateTime(v string) *GetTrustedOriginResponseBodyTrustedOrigin {
	s.CreateTime = &v
	return s
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) SetInstanceId(v string) *GetTrustedOriginResponseBodyTrustedOrigin {
	s.InstanceId = &v
	return s
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) SetOrigin(v string) *GetTrustedOriginResponseBodyTrustedOrigin {
	s.Origin = &v
	return s
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) SetStatus(v string) *GetTrustedOriginResponseBodyTrustedOrigin {
	s.Status = &v
	return s
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) SetTrustOriginName(v string) *GetTrustedOriginResponseBodyTrustedOrigin {
	s.TrustOriginName = &v
	return s
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) SetTrustedOriginId(v string) *GetTrustedOriginResponseBodyTrustedOrigin {
	s.TrustedOriginId = &v
	return s
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) SetTrustedOriginScene(v []*string) *GetTrustedOriginResponseBodyTrustedOrigin {
	s.TrustedOriginScene = v
	return s
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) SetUpdateTime(v string) *GetTrustedOriginResponseBodyTrustedOrigin {
	s.UpdateTime = &v
	return s
}

func (s *GetTrustedOriginResponseBodyTrustedOrigin) Validate() error {
	return dara.Validate(s)
}
