// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSignatureLibVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSignatureLibVersionResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeSignatureLibVersionResponseBody
	GetTotalCount() *int32
	SetVersion(v []*DescribeSignatureLibVersionResponseBodyVersion) *DescribeSignatureLibVersionResponseBody
	GetVersion() []*DescribeSignatureLibVersionResponseBodyVersion
}

type DescribeSignatureLibVersionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9C50C2A9-4BBB-5504-8ADA-C41A79B8C946
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 132
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The version information.
	Version []*DescribeSignatureLibVersionResponseBodyVersion `json:"Version,omitempty" xml:"Version,omitempty" type:"Repeated"`
}

func (s DescribeSignatureLibVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignatureLibVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSignatureLibVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSignatureLibVersionResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSignatureLibVersionResponseBody) GetVersion() []*DescribeSignatureLibVersionResponseBodyVersion {
	return s.Version
}

func (s *DescribeSignatureLibVersionResponseBody) SetRequestId(v string) *DescribeSignatureLibVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSignatureLibVersionResponseBody) SetTotalCount(v int32) *DescribeSignatureLibVersionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSignatureLibVersionResponseBody) SetVersion(v []*DescribeSignatureLibVersionResponseBodyVersion) *DescribeSignatureLibVersionResponseBody {
	s.Version = v
	return s
}

func (s *DescribeSignatureLibVersionResponseBody) Validate() error {
	if s.Version != nil {
		for _, item := range s.Version {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSignatureLibVersionResponseBodyVersion struct {
	// The type.
	//
	// Valid values:
	//
	// 	- ips
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Basic Rules and Virtual Patching
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- intelligence
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Threat Intelligence
	//
	//     <!-- -->
	//
	// example:
	//
	// ips
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Update time.
	//
	// example:
	//
	// 1741067915
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version number.
	//
	// example:
	//
	// IPS-2307-02
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeSignatureLibVersionResponseBodyVersion) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignatureLibVersionResponseBodyVersion) GoString() string {
	return s.String()
}

func (s *DescribeSignatureLibVersionResponseBodyVersion) GetType() *string {
	return s.Type
}

func (s *DescribeSignatureLibVersionResponseBodyVersion) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeSignatureLibVersionResponseBodyVersion) GetVersion() *string {
	return s.Version
}

func (s *DescribeSignatureLibVersionResponseBodyVersion) SetType(v string) *DescribeSignatureLibVersionResponseBodyVersion {
	s.Type = &v
	return s
}

func (s *DescribeSignatureLibVersionResponseBodyVersion) SetUpdateTime(v int64) *DescribeSignatureLibVersionResponseBodyVersion {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSignatureLibVersionResponseBodyVersion) SetVersion(v string) *DescribeSignatureLibVersionResponseBodyVersion {
	s.Version = &v
	return s
}

func (s *DescribeSignatureLibVersionResponseBodyVersion) Validate() error {
	return dara.Validate(s)
}
