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
	// The user-generated request token. This token is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the original production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance ID/Name column.
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
