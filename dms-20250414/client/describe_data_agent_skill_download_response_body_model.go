// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAgentSkillDownloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDataAgentSkillDownloadResponseBodyData) *DescribeDataAgentSkillDownloadResponseBody
	GetData() *DescribeDataAgentSkillDownloadResponseBodyData
	SetErrorCode(v string) *DescribeDataAgentSkillDownloadResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeDataAgentSkillDownloadResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeDataAgentSkillDownloadResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDataAgentSkillDownloadResponseBody
	GetSuccess() *bool
}

type DescribeDataAgentSkillDownloadResponseBody struct {
	// The response struct.
	Data *DescribeDataAgentSkillDownloadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when a request exception occurs.
	//
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDataAgentSkillDownloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSkillDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSkillDownloadResponseBody) GetData() *DescribeDataAgentSkillDownloadResponseBodyData {
	return s.Data
}

func (s *DescribeDataAgentSkillDownloadResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDataAgentSkillDownloadResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDataAgentSkillDownloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataAgentSkillDownloadResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDataAgentSkillDownloadResponseBody) SetData(v *DescribeDataAgentSkillDownloadResponseBodyData) *DescribeDataAgentSkillDownloadResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDataAgentSkillDownloadResponseBody) SetErrorCode(v string) *DescribeDataAgentSkillDownloadResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDataAgentSkillDownloadResponseBody) SetErrorMessage(v string) *DescribeDataAgentSkillDownloadResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDataAgentSkillDownloadResponseBody) SetRequestId(v string) *DescribeDataAgentSkillDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataAgentSkillDownloadResponseBody) SetSuccess(v bool) *DescribeDataAgentSkillDownloadResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDataAgentSkillDownloadResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataAgentSkillDownloadResponseBodyData struct {
	// The public download link.
	//
	// example:
	//
	// https://*******.oss-cn-hangzhou.aliyuncs.com/data_agent/skill_download/1673826******66/ski-1giw1afhxxl6********ygcir/*******.zip?Expires=1793935883&OSSAccessKeyId=STS.NXy5M8bN5wuNt*******NxDW&Signature=%2FrU3FTVCAdtJJvZQvrftw6AWQug%3D&response-content-disposition=attachment&security-token=CAISiQR1q6Ft5B2yfSjIr5vMfveMj5EUwLeldmzA1lA6ZLlBob3vtTz2IHhMdXlrCOgYt%2F8xnG1V6f8flrJ%2FToQAX0HfatZq5ZkS9AqnaoXM%2Fte496IFg5D9pf1flOynjqHoeOzcYI734p7PAgm2FEYRrJL%2BcTK9JRPHVbSClZ9gaPkOQwC8dkAoLdxKJwxk2qh4UHzKLqSCPwLNiGjdB1YSZWMe71ly8qOi2MaRxwPDhVnhsI8vqp%2F2P4KvYrueVfWAOt6%2Bx75xbbGTknwSuRRP7%2BJ6hbBNvD7Is9eBRlJXpg6GOuzb8tRuawNweqENEqFJquS7jv55vujkm4PozRJKNO0SCnqAH9D6npWZQLr0bI9nK%2FHgOHWIi9%2BOKcD8vQM0cXZFKwRYKtQ8LWBhBhctUHbKMqPi9lTNfUK%2FQrOGlagt0pZ801zfpIXbfAnUGe%2FJi3dBY9kublsKT0dOhTG5K%2FJfLlEcKAg5WeevIN8tPEwEm8TBphbPUyBs9HZTssDlaune0qJlMt2nD8IdgNJDPc8Y6zZ3EA2qUcGnjkYFaWppRK1Oyq7mNISk6bmKzOyPaPTBDv0Xb%2BrupowlAUBnIFEnzk2ZjrVLATLwcBQLRCBzwt1KpKo%2FNH0u7BAEADeclSYSR4GYQXr9nsHtPA6ksAQWlPeWMacBiy9GOsPAqYv0GlbRDSX9P8FR8BqAARLcIr26y%2BphYuiqaegPh83GYEiMKKQS66fO%2FbR98Z9kD2BZrL%2FN3hztB48fQZyNeNE9AMmLzxuyk%2BHfEozrUOP7nvr%2F73NxdW5ydDWgil66VEwx5Grjy43Q3OZPIPUrn7t4WFMWd8oxvqgo47IApLHY3G5akK6pzYHmVrskTrtjIAA%3D
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// The internal network download link.
	//
	// example:
	//
	// https://*******.oss-cn-hangzhou-internal.aliyuncs.com/data_agent/skill_download/1673826******66/ski-1giw1afhxxl6********ygcir/*******.zip?Expires=1793935883&OSSAccessKeyId=STS.NXy5M8bN5wuNt*******NxDW&Signature=%2FrU3FTVCAdtJJvZQvrftw6AWQug%3D&response-content-disposition=attachment&security-token=CAISiQR1q6Ft5B2yfSjIr5vMfveMj5EUwLeldmzA1lA6ZLlBob3vtTz2IHhMdXlrCOgYt%2F8xnG1V6f8flrJ%2FToQAX0HfatZq5ZkS9AqnaoXM%2Fte496IFg5D9pf1flOynjqHoeOzcYI734p7PAgm2FEYRrJL%2BcTK9JRPHVbSClZ9gaPkOQwC8dkAoLdxKJwxk2qh4UHzKLqSCPwLNiGjdB1YSZWMe71ly8qOi2MaRxwPDhVnhsI8vqp%2F2P4KvYrueVfWAOt6%2Bx75xbbGTknwSuRRP7%2BJ6hbBNvD7Is9eBRlJXpg6GOuzb8tRuawNweqENEqFJquS7jv55vujkm4PozRJKNO0SCnqAH9D6npWZQLr0bI9nK%2FHgOHWIi9%2BOKcD8vQM0cXZFKwRYKtQ8LWBhBhctUHbKMqPi9lTNfUK%2FQrOGlagt0pZ801zfpIXbfAnUGe%2FJi3dBY9kublsKT0dOhTG5K%2FJfLlEcKAg5WeevIN8tPEwEm8TBphbPUyBs9HZTssDlaune0qJlMt2nD8IdgNJDPc8Y6zZ3EA2qUcGnjkYFaWppRK1Oyq7mNISk6bmKzOyPaPTBDv0Xb%2BrupowlAUBnIFEnzk2ZjrVLATLwcBQLRCBzwt1KpKo%2FNH0u7BAEADeclSYSR4GYQXr9nsHtPA6ksAQWlPeWMacBiy9GOsPAqYv0GlbRDSX9P8FR8BqAARLcIr26y%2BphYuiqaegPh83GYEiMKKQS66fO%2FbR98Z9kD2BZrL%2FN3hztB48fQZyNeNE9AMmLzxuyk%2BHfEozrUOP7nvr%2F73NxdW5ydDWgil66VEwx5Grjy43Q3OZPIPUrn7t4WFMWd8oxvqgo47IApLHY3G5akK6pzYHmVrskTrtjIAA%3D
	IntranetDownloadLink *string `json:"IntranetDownloadLink,omitempty" xml:"IntranetDownloadLink,omitempty"`
	// The skill ID.
	//
	// example:
	//
	// ski-04pomiln*************j0
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The skill name.
	//
	// example:
	//
	// data-query-skill
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s DescribeDataAgentSkillDownloadResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSkillDownloadResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSkillDownloadResponseBodyData) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *DescribeDataAgentSkillDownloadResponseBodyData) GetIntranetDownloadLink() *string {
	return s.IntranetDownloadLink
}

func (s *DescribeDataAgentSkillDownloadResponseBodyData) GetSkillId() *string {
	return s.SkillId
}

func (s *DescribeDataAgentSkillDownloadResponseBodyData) GetSkillName() *string {
	return s.SkillName
}

func (s *DescribeDataAgentSkillDownloadResponseBodyData) SetDownloadLink(v string) *DescribeDataAgentSkillDownloadResponseBodyData {
	s.DownloadLink = &v
	return s
}

func (s *DescribeDataAgentSkillDownloadResponseBodyData) SetIntranetDownloadLink(v string) *DescribeDataAgentSkillDownloadResponseBodyData {
	s.IntranetDownloadLink = &v
	return s
}

func (s *DescribeDataAgentSkillDownloadResponseBodyData) SetSkillId(v string) *DescribeDataAgentSkillDownloadResponseBodyData {
	s.SkillId = &v
	return s
}

func (s *DescribeDataAgentSkillDownloadResponseBodyData) SetSkillName(v string) *DescribeDataAgentSkillDownloadResponseBodyData {
	s.SkillName = &v
	return s
}

func (s *DescribeDataAgentSkillDownloadResponseBodyData) Validate() error {
	return dara.Validate(s)
}
