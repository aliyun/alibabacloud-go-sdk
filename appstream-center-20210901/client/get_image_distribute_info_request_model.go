// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDistributeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *GetImageDistributeInfoRequest
	GetImageId() *string
	SetVersionId(v string) *GetImageDistributeInfoRequest
	GetVersionId() *string
}

type GetImageDistributeInfoRequest struct {
	// The ID of the image whose region distribution information you want to query.
	//
	// example:
	//
	// imgc-bp1example****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the image version. If this parameter is not specified, the online version of the image is queried.
	//
	// example:
	//
	// iv-xxx
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetImageDistributeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageDistributeInfoRequest) GoString() string {
	return s.String()
}

func (s *GetImageDistributeInfoRequest) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageDistributeInfoRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetImageDistributeInfoRequest) SetImageId(v string) *GetImageDistributeInfoRequest {
	s.ImageId = &v
	return s
}

func (s *GetImageDistributeInfoRequest) SetVersionId(v string) *GetImageDistributeInfoRequest {
	s.VersionId = &v
	return s
}

func (s *GetImageDistributeInfoRequest) Validate() error {
	return dara.Validate(s)
}
