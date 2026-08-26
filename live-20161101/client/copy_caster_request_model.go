// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCasterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterName(v string) *CopyCasterRequest
	GetCasterName() *string
	SetClientToken(v string) *CopyCasterRequest
	GetClientToken() *string
	SetOwnerId(v int64) *CopyCasterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CopyCasterRequest
	GetRegionId() *string
	SetSrcCasterId(v string) *CopyCasterRequest
	GetSrcCasterId() *string
}

type CopyCasterRequest struct {
	// The name of the new production studio.
	//
	// This parameter is required.
	//
	// example:
	//
	// caster001
	CasterName *string `json:"CasterName,omitempty" xml:"CasterName,omitempty"`
	// A client-generated token that ensures the idempotence of the request.
	//
	// Generate a unique value for this parameter for each request. The token can contain a maximum of 64 ASCII characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the production studio to copy.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value that is returned.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to **ApsaraVideo Live console*	- > **Production Studio*	- > **Cloud Production Studio*	- to view the production studio name.
	//
	// > The name of a production studio on the Cloud Production Studio page is its production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	SrcCasterId *string `json:"SrcCasterId,omitempty" xml:"SrcCasterId,omitempty"`
}

func (s CopyCasterRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyCasterRequest) GoString() string {
	return s.String()
}

func (s *CopyCasterRequest) GetCasterName() *string {
	return s.CasterName
}

func (s *CopyCasterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CopyCasterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CopyCasterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CopyCasterRequest) GetSrcCasterId() *string {
	return s.SrcCasterId
}

func (s *CopyCasterRequest) SetCasterName(v string) *CopyCasterRequest {
	s.CasterName = &v
	return s
}

func (s *CopyCasterRequest) SetClientToken(v string) *CopyCasterRequest {
	s.ClientToken = &v
	return s
}

func (s *CopyCasterRequest) SetOwnerId(v int64) *CopyCasterRequest {
	s.OwnerId = &v
	return s
}

func (s *CopyCasterRequest) SetRegionId(v string) *CopyCasterRequest {
	s.RegionId = &v
	return s
}

func (s *CopyCasterRequest) SetSrcCasterId(v string) *CopyCasterRequest {
	s.SrcCasterId = &v
	return s
}

func (s *CopyCasterRequest) Validate() error {
	return dara.Validate(s)
}
