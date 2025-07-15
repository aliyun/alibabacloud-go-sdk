// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAndroidInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *StartAndroidInstanceRequest
	GetAndroidInstanceIds() []*string
	SetSaleMode(v string) *StartAndroidInstanceRequest
	GetSaleMode() *string
}

type StartAndroidInstanceRequest struct {
	// List of instances.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	SaleMode           *string   `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
}

func (s StartAndroidInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAndroidInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartAndroidInstanceRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *StartAndroidInstanceRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *StartAndroidInstanceRequest) SetAndroidInstanceIds(v []*string) *StartAndroidInstanceRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *StartAndroidInstanceRequest) SetSaleMode(v string) *StartAndroidInstanceRequest {
	s.SaleMode = &v
	return s
}

func (s *StartAndroidInstanceRequest) Validate() error {
	return dara.Validate(s)
}
