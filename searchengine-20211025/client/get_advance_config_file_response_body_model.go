// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvanceConfigFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAdvanceConfigFileResponseBody
	GetRequestId() *string
	SetResult(v *GetAdvanceConfigFileResponseBodyResult) *GetAdvanceConfigFileResponseBody
	GetResult() *GetAdvanceConfigFileResponseBodyResult
}

type GetAdvanceConfigFileResponseBody struct {
	// id of request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	Result *GetAdvanceConfigFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAdvanceConfigFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAdvanceConfigFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAdvanceConfigFileResponseBody) GetResult() *GetAdvanceConfigFileResponseBodyResult {
	return s.Result
}

func (s *GetAdvanceConfigFileResponseBody) SetRequestId(v string) *GetAdvanceConfigFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdvanceConfigFileResponseBody) SetResult(v *GetAdvanceConfigFileResponseBodyResult) *GetAdvanceConfigFileResponseBody {
	s.Result = v
	return s
}

func (s *GetAdvanceConfigFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAdvanceConfigFileResponseBodyResult struct {
	// The file content.
	//
	// example:
	//
	// {"summarys":{"parameter":{"file_compressor":"zstd"},"summary_fields":["id"]},"file_compress":[{"name":"file_compressor","type":"zstd"},{"name":"no_compressor","type":""}],"indexs":[{"index_fields":"name","index_name":"ids","index_type":"STRING"},{"has_primary_key_attribute":true,"index_fields":"id","is_primary_key_sorted":false,"index_name":"id","index_type":"PRIMARYKEY64"}],"attributes":[{"file_compress":"no_compressor","field_name":"id"}],"fields":[{"user_defined_param":{},"compress_type":"uniq","field_type":"STRING","field_name":"id"},{"compress_type":"uniq","field_type":"STRING","field_name":"name"}],"table_name":"api"}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s GetAdvanceConfigFileResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetAdvanceConfigFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigFileResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *GetAdvanceConfigFileResponseBodyResult) SetContent(v string) *GetAdvanceConfigFileResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetAdvanceConfigFileResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
