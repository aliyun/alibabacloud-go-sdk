// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDatasetDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataArray(v []map[string]interface{}) *AddDatasetDataRequest
	GetDataArray() []map[string]interface{}
	SetClientToken(v string) *AddDatasetDataRequest
	GetClientToken() *string
}

type AddDatasetDataRequest struct {
	DataArray []map[string]interface{} `json:"dataArray,omitempty" xml:"dataArray,omitempty" type:"Repeated"`
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s AddDatasetDataRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDataRequest) GoString() string {
	return s.String()
}

func (s *AddDatasetDataRequest) GetDataArray() []map[string]interface{} {
	return s.DataArray
}

func (s *AddDatasetDataRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddDatasetDataRequest) SetDataArray(v []map[string]interface{}) *AddDatasetDataRequest {
	s.DataArray = v
	return s
}

func (s *AddDatasetDataRequest) SetClientToken(v string) *AddDatasetDataRequest {
	s.ClientToken = &v
	return s
}

func (s *AddDatasetDataRequest) Validate() error {
	return dara.Validate(s)
}
