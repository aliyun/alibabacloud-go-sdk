// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDefineRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteUserDefineRegionResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteUserDefineRegionResponseBody
	GetMessage() *string
	SetRegionDefine(v *DeleteUserDefineRegionResponseBodyRegionDefine) *DeleteUserDefineRegionResponseBody
	GetRegionDefine() *DeleteUserDefineRegionResponseBodyRegionDefine
	SetRequestId(v string) *DeleteUserDefineRegionResponseBody
	GetRequestId() *string
}

type DeleteUserDefineRegionResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The custom namespace.
	RegionDefine *DeleteUserDefineRegionResponseBodyRegionDefine `json:"RegionDefine,omitempty" xml:"RegionDefine,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1234-1sda-321d-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserDefineRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDefineRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserDefineRegionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteUserDefineRegionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteUserDefineRegionResponseBody) GetRegionDefine() *DeleteUserDefineRegionResponseBodyRegionDefine {
	return s.RegionDefine
}

func (s *DeleteUserDefineRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserDefineRegionResponseBody) SetCode(v int32) *DeleteUserDefineRegionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBody) SetMessage(v string) *DeleteUserDefineRegionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBody) SetRegionDefine(v *DeleteUserDefineRegionResponseBodyRegionDefine) *DeleteUserDefineRegionResponseBody {
	s.RegionDefine = v
	return s
}

func (s *DeleteUserDefineRegionResponseBody) SetRequestId(v string) *DeleteUserDefineRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBody) Validate() error {
	if s.RegionDefine != nil {
		if err := s.RegionDefine.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteUserDefineRegionResponseBodyRegionDefine struct {
	// The ID of the region to which the custom namespace belongs.
	//
	// example:
	//
	// cn-beijing
	BelongRegion *string `json:"BelongRegion,omitempty" xml:"BelongRegion,omitempty"`
	// The description of the custom namespace.
	//
	// example:
	//
	// logic region
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identifier of the custom namespace.
	//
	// example:
	//
	// 8848
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the custom namespace. The ID cannot be changed after the custom namespace is created. The format is `region ID:custom namespace ID`.
	//
	// example:
	//
	// cn-beijing:test
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the custom namespace.
	//
	// example:
	//
	// test_region
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The ID of the Alibaba Cloud account to which the custom namespace belongs.
	//
	// example:
	//
	// 11727****22398
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserDefineRegionResponseBodyRegionDefine) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDefineRegionResponseBodyRegionDefine) GoString() string {
	return s.String()
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) GetBelongRegion() *string {
	return s.BelongRegion
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) GetDescription() *string {
	return s.Description
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) GetId() *int64 {
	return s.Id
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) GetRegionName() *string {
	return s.RegionName
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) GetUserId() *string {
	return s.UserId
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) SetBelongRegion(v string) *DeleteUserDefineRegionResponseBodyRegionDefine {
	s.BelongRegion = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) SetDescription(v string) *DeleteUserDefineRegionResponseBodyRegionDefine {
	s.Description = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) SetId(v int64) *DeleteUserDefineRegionResponseBodyRegionDefine {
	s.Id = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) SetRegionId(v string) *DeleteUserDefineRegionResponseBodyRegionDefine {
	s.RegionId = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) SetRegionName(v string) *DeleteUserDefineRegionResponseBodyRegionDefine {
	s.RegionName = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) SetUserId(v string) *DeleteUserDefineRegionResponseBodyRegionDefine {
	s.UserId = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) Validate() error {
	return dara.Validate(s)
}
