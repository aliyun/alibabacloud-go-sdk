// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedDataKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataKeyName(v string) *GetManagedDataKeyRequest
	GetDataKeyName() *string
	SetDataKeyVersionId(v string) *GetManagedDataKeyRequest
	GetDataKeyVersionId() *string
	SetUseLatest(v bool) *GetManagedDataKeyRequest
	GetUseLatest() *bool
}

type GetManagedDataKeyRequest struct {
	// The name of the managed data key (DK). This parameter is required.
	//
	// example:
	//
	// example-data-key
	DataKeyName *string `json:"DataKeyName,omitempty" xml:"DataKeyName,omitempty"`
	// The version number of the managed data key (DK). This parameter is optional. If you set this parameter to a specific version number, the plaintext of the specified version of the managed data key (DK) is returned.
	//
	// example:
	//
	// xH6OPUmz
	DataKeyVersionId *string `json:"DataKeyVersionId,omitempty" xml:"DataKeyVersionId,omitempty"`
	// Specifies whether to use the latest version of the managed data key (DK) when no version number is provided. Valid values:
	//
	// - true: Returns the latest version of the managed data key (DK).
	//
	// - false: Returns the first version of the managed data key (DK).
	//
	// Default value: false.
	//
	// example:
	//
	// true
	UseLatest *bool `json:"UseLatest,omitempty" xml:"UseLatest,omitempty"`
}

func (s GetManagedDataKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetManagedDataKeyRequest) GoString() string {
	return s.String()
}

func (s *GetManagedDataKeyRequest) GetDataKeyName() *string {
	return s.DataKeyName
}

func (s *GetManagedDataKeyRequest) GetDataKeyVersionId() *string {
	return s.DataKeyVersionId
}

func (s *GetManagedDataKeyRequest) GetUseLatest() *bool {
	return s.UseLatest
}

func (s *GetManagedDataKeyRequest) SetDataKeyName(v string) *GetManagedDataKeyRequest {
	s.DataKeyName = &v
	return s
}

func (s *GetManagedDataKeyRequest) SetDataKeyVersionId(v string) *GetManagedDataKeyRequest {
	s.DataKeyVersionId = &v
	return s
}

func (s *GetManagedDataKeyRequest) SetUseLatest(v bool) *GetManagedDataKeyRequest {
	s.UseLatest = &v
	return s
}

func (s *GetManagedDataKeyRequest) Validate() error {
	return dara.Validate(s)
}
